Now that [version 2](../v2/ReadMe.md) is completed, for all intents and purposes, we have a complete notification system!  However this is still far from complete.  We currently have a few blatant issues with our system that should be addressed:

- **Insufficient Admin Authentication**: The admin portal is vaguely "password protected," which honestly just isn't good enough. A proper authentication and authorization system, is needed to manage admin users, enforce strong password policies, and ensure only authorized individuals can send notifications.

- **Vulnerable API**: Our public-facing endpoints are open to the internet without any protection. This makes them vulnerable to bots and attackers who could spam the service, subscribe users without their consent, or trigger a flood of Lambda invocations.

- **Subscription Management**: Users can only subscribe or unsubscribe. They cannot choose what types of notifications they want to receive.  On a similar note, Admins can only send to all subscribers rather than allowing for targeted notifications.

As such, we will be expanding the scope of our system to address these key concerns for V3.

> As an aside, I will note that code example will become less frequent throughout this project as the primary purpose is as a system design exercise rather than a coding exercise.