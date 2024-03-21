variable "aws_region" {
  type    = string
  default = "us-east-1"
}

variable "vpc_id" {
  type    = string
  default = ""
}

variable "subnet_a" {
  type    = string
  default = ""
}

variable "subnet_b" {
  type    = string
  default = ""
}

variable "ecr_image" {
  description = "ECR Image"
  type        = string
  sensitive   = true
  default     = ""
}

variable "execution_role_ecs" {
  description = "Execution role ECS"
  type        = string
  sensitive   = true
  default     = ""
}

variable "desired_tasks" {
  description = "Mininum executing tasks"
  type        = number
  default     = 1
}

variable "ecs_cluster" {
  description = "Cluster ECS ARN"
  type        = string
  sensitive   = true
  default     = ""
}

variable "sg_cluster_ecs" {
  description = "Cluster ECS Security group"
  type        = string
  default     = ""
}

variable "lb_arn" {
  description = "Load Balancer ARN"
  type        = string
  sensitive   = true
  default     = ""
}

variable "alb_fastfood_listener_arn" {
  description = "Default Listener ALB"
  type        = string
  sensitive   = true
  default     = ""
}

variable "db_url" {
  description = "Hackathon DB URL"
  type        = string
  sensitive   = true
  default     = ""
}

variable "sns_arn" {
  description = "SNS Hackathon URL"
  type        = string
  sensitive   = true
  default     = ""
}
