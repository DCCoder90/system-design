Now that [version 2](../v2/ReadMe.md) is completed, for all intents and purposes, we have a complete notification system!  However this is still far from complete.  We currently have a few blatant issues with our system that should be addressed:

- **Insufficient Admin Authentication**: The admin portal is vaguely "password protected," which honestly just isn't good enough. A proper authentication and authorization system, is needed to manage admin users, enforce strong password policies, and ensure only authorized individuals can send notifications.

- **Vulnerable API**: Our public-facing endpoints are open to the internet without any protection. This makes them vulnerable to bots and attackers who could spam the service, subscribe users without their consent, or trigger a flood of Lambda invocations.

- **Subscription Management**: Users can only subscribe or unsubscribe. They cannot choose what types of notifications they want to receive.  On a similar note, Admins can only send to all subscribers rather than allowing for targeted notifications.

As such, we will be expanding the scope of our system to address these key concerns for V3.

> As an aside, I will note that code example will become less frequent throughout this project as the primary purpose is as a system design exercise rather than a coding exercise.


> It's been about a week since my last update, but such is life at times.  Between job hunting, work at my job, family, birthdays, etc.  Things got a little hectic here for a bit.  But I am still working on this one.


### Admin Authentication

For administrator authentication, our approach will focus on using [Amazon Cognito](https://docs.aws.amazon.com/cognito/). We'll establish a [Cognito user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools.html) to serve as our secure, managed directory for all admin accounts. This choice offloads the heavy lifting of user management, as Cognito natively handles the complete lifecycle: admin sign-up, sign-in, password policies, and MFA. It also provides a hosted UI, which we can integrate directly with the Admin Portal to manage the login process.

The authentication itself is fairly straightforward. When an admin accesses the Admin Portal, they'll be redirected to the Cognito-hosted UI. Upon successful authentication, Cognito will issue a JWT. The portal's frontend application will then be responsible for including this JWT in the Authorization header for all subsequent API calls to protected routes, such as the `/send` or new `/topics` endpoints.

On the backend, thsi will be enforced by API Gateway. We will attach a [Cognito Authorizer](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html) to all admin-only routes. This authorizer automatically intercepts incoming requests before they reach our Lambda functions. It checks the JWT, validates the signature, and confirms the user is real. This integration means API Gateway handles all security checks; if a token is missing, expired, or invalid, the request is immediately rejected with a `401 Unauthorized` error.

### Securing the API

Our public-facing endpoints, `/subscribe` and `/unsubscribe`, are vulnerable to bots and other abuse. To lock this down, we'll implement [AWS WAF](https://aws.amazon.com/waf/) as our first line of defense. We can attach a [WAF WebACL](https://docs.aws.amazon.com/waf/latest/developerguide/web-acl.html) directly to our API Gateway.  This will let WAF inspect and filter traffic at this stage before reaching our lambdas.

This gives us two immediate layers of protection. First, we'll enable the [AWSManagedRulesCommonRuleSet](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-baseline.html). This is a pre-configured ruleset maintained by AWS that automatically blocks traffic from known bots, scanners, and malicious IP addresses. Second, we'll add a custom rate-based rule aimed at `/subscribe`. This will help prevent simple spam scripts by limiting requests.

While using the WAF is a great first step, we also need to add a little extra protection to the front-end.  To solve this we will implement a CAPTCHA (although that may not be enough in today's AI age) on the public facing web portal. The "subscribe" button will remain inactive until the CAPTCHA passes. 

### Subscription Management
This is going to be the largest change since starting this project.  This will move the notification system from a single to a flexible, multi-topic model. This will allow us to create distinct notification categories, such as "News", "Events", "Promotions", etc.  

First, we'll introduce a new DynamoDB table, which we will call `TopicsTable`. This will allow us to map our topics to their full SNS `topic_arn`. Next we'll modify the existing `UsersTable`. We're going to add a new `subscriptions` attribute, which will be a StringSet. This set will hold the names of all the topics a specific user has subscribed to.
