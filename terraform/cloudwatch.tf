resource "aws_cloudwatch_log_group" "ecs_backend_log_group" {
  name              = "/ecs/ecs-backend-task"
  retention_in_days = 7
}

resource "aws_cloudwatch_log_group" "ecs_frontend_log_group" {
  name              = "/ecs/ecs-frontend-task"
  retention_in_days = 7
}
