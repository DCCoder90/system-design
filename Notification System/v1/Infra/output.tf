output "sns_topic_arn" {
  value       = aws_sns_topic.notification_topic.arn
  description = "The ARN of the SNS topic."
}

output "user_access_key_id" {
  value       = aws_iam_access_key.user_key.id
  description = "The access key ID for the IAM user."
}

output "user_secret_access_key" {
  value       = aws_iam_access_key.user_key.secret
  description = "The secret access key for the IAM user."
  sensitive   = true
}