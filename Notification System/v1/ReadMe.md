The first version is a Minimum Viable Product (MVP). This should only cover the extreme basics of the Notification System, as such we only want to focus on a product that can generate and send notifications.  We are not concerned with anything beyond that scope for the moment.  This would typically be done as a proof of concept, to simply prove we can do it without spending too much time.  The primary purpose here is to balance budget and level of work.



\## Cloud Provider

To begin, we'll start looking at cloud providers to minimize effort.  AWS is a good first choice, being an industry leader and well established in the cloud space.



Looking at AWS's offerings we can see a product called SNS or (Simple Notification Services).  Through AWS SNS we can send notifications through the following methods:



* Amazon SQS
* Lambda
* HTTP(S) endpoints
* Email
* Mobile push notifications
* Mobile text messages (SMS)
* Amazon Data Firehose
* Service providers (For example, Datadog, MongoDB, Splunk)



https://docs.aws.amazon.com/sns/latest/dg/welcome.html



Seeing as our core focus is SMS, Email, and push notifications it appears as though SNS covers all of these topics and then some, this allows us to get a quick start with the ability to expand in the future if we so wish.  SNS also has reasonable pricing amounts as well.  This can be viewed on their \[pricing page](https://aws.amazon.com/sns/pricing/), however I will include a table below as well.



\### API Requests

* Standard topic requests include publishes, batch publishes, topic owner operations, and subscription owner operations
* First 1 million Amazon SNS requests per month are free, $0.50 per 1 million requests thereafter

\*\*\*Note:\*\*\* Each 64KB chunk of published data is billed as 1 request. For example, a single publish with a 256KB payload is billed as four requests.



\### Notification Deliveries

|Endpoint Type|Free Tier|Price|

|---|---|---|

|Mobile Push Notifications|1 million notifications|$0.50 per million notifications|

|Email/Email-JSON|1,000 notifications|$2.00 per 100,000 notifications|

|HTTP/s|100,000 notifications|$0.60 per million notifications|

|Simple Queue Service (SQS)|No charge for deliveries to SQS Queues. Standard SQS pricing applies. Data transfer charges apply between Amazon SNS and Amazon SQS.|---|

|AWS Lambda|No charge for deliveries to Lambda. Standard Lambda pricing applies. Data transfer charges apply between Amazon SNS and Lambda.|---|

|Amazon Data Firehose|Standard Amazon Data Firehose pricing applies. Data transfer charges apply between SNS and Amazon Data Firehose.|$0.19 per million notifications|



\## Infra



I'll start by using \[Terraform](https://developer.hashicorp.com/terraform) to manage the infrastructure.  This will allow us to more easily reproduce the exercise as well as quickly destroy all resources that we create when we are done with them since this is just a proof of concept/MVP.



The infrastructure is being stored int the \[Infra directory](./Infra).   This includes some basic terraform to simply create a new user with access keys that our application can use as well as create the SNS topic and relevant IAM policies to ensure that we can use SNS.



