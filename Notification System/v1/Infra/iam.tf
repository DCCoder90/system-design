resource "aws_iam_user" "user" {
  name = "sns-notifier-user"
}

data "aws_iam_policy_document" "policy_doc" {
  statement {
    effect = "Allow"
    actions = [
      "sns:Publish",
      "sns:Subscribe"
    ]
    resources = [
      aws_sns_topic.notification_topic.arn
    ]
  }
}

resource "aws_iam_policy" "policy" {
  name   = "Notification-System-Policy"
  policy = data.aws_iam_policy_document.policy_doc.json
}

resource "aws_iam_user_policy_attachment" "app_policy_attach" {
  user       = aws_iam_user.user.name
  policy_arn = aws_iam_policy.policy.arn
}

resource "aws_iam_access_key" "user_key" {
  user = aws_iam_user.user.name
}