package console

// ServiceMap maps CLI flags to AWS console URL paths.
// e.g. passing in `-s ec2` will open the console at the ec2/v2 URL.
var ServiceMap = map[string]string{
	"":               "console",
	"acm":            "acm",
	"aos":            "aos",
	"apigateway":     "apigateway",
	"apigw":          "apigateway",
	"appsync":        "appsync",
	"athena":         "athena",
	"bedrock":        "bedrock",
	"c9":             "cloud9",
	"ce":             "cost-management",
	"cf":             "cloudfront",
	"cfn":            "cloudformation",
	"cloudformation": "cloudformation",
	"cloudmap":       "cloudmap",
	"cloudwatch":     "cloudwatch",
	"codeartifact":   "codesuite/codeartifact",
	"codecommit":     "codesuite/codecommit",
	"codedeploy":     "codedeploy",
	"codepipeline":   "codepipeline",
	"codesuite":      "codesuite",
	"config":         "config",
	"ct":             "cloudtrail",
	"cw":             "cloudwatch",
	"ddb":            "dynamodbv2",
	"dms":            "dms/v2",
	"dx":             "directconnect/v2",
	"dynamodb":       "dynamodbv2",
	"eb":             "elasticbeanstalk",
	"ebs":            "elasticbeanstalk",
	"ec2":            "ec2/v2",
	"ecr":            "ecr",
	"ecs":            "ecs",
	"eks":            "eks",
	"eventbridge":    "events",
	"events":         "events",
	"gd":             "guardduty",
	"grafana":        "grafana",
	"iam":            "iamv2",
	"kms":            "kms",
	"l":              "lambda",
	"lambda":         "lambda",
	"mwaa":           "mwaa",
	"param":          "systems-manager/parameters",
	"r53":            "route53/v2",
	"ram":            "ram",
	"rds":            "rds",
	"redshift":       "redshiftv2",
	"route53":        "route53/v2",
	"s3":             "s3",
	"sagemaker":      "sagemaker",
	"scrh":           "securityhub",
	"scrm":           "secretsmanager",
	"secretsmanager": "secretsmanager",
	"securityhub":    "securityhub",
	"ses":            "ses",
	"sfn":            "states",
	"sm":             "secretsmanager",
	"sns":            "sns",
	"ssm":            "systems-manager",
	"sso":            "singlesignon",
	"states":         "states",
	"stepfn":         "states",
	"tra":            "trustedadvisor",
	"trustedadvisor": "trustedadvisor",
	"vpc":            "vpc",
	"waf":            "wafv2",
	"wafv2":          "wafv2/homev2",
}

var globalServiceMap = map[string]bool{
	"dx":             true,
	"iam":            true,
	"r53":            true,
	"route53":        true,
	"trustedadvisor": true,
}
