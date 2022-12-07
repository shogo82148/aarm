local branch = "main";

{
  "ServiceName": "my-app-runner-test2",
  "SourceConfiguration": {
    "AuthenticationConfiguration": {
      "ConnectionArn": "arn:aws:apprunner:ap-northeast-1:445285296882:connection/shogo82148/a52503ce07314ca99c4e13c47c25877f"
    },
    "AutoDeploymentsEnabled": true,
    "CodeRepository": {
      "CodeConfiguration": {
        "CodeConfigurationValues": {
          "BuildCommand": "go build -o myapp .",
          "Port": "8000",
          "Runtime": "GO_1",
          "StartCommand": "./myapp"
        },
        "ConfigurationSource": "API"
      },
      "RepositoryUrl": "https://github.com/shogo82148/aws-app-runner-test",
      "SourceCodeVersion": {
        "Type": "BRANCH",
        "Value": branch,
      }
    }
  },
  "AutoScalingConfigurationArn": "arn:aws:apprunner:ap-northeast-1:445285296882:autoscalingconfiguration/DefaultConfiguration/1/00000000000000000000000000000001",
  "HealthCheckConfiguration": {
    "HealthyThreshold": 1,
    "Interval": 10,
    "Path": "/",
    "Protocol": "TCP",
    "Timeout": 5,
    "UnhealthyThreshold": 5
  },
  "InstanceConfiguration": {
    "Cpu": "1024",
    "Memory": "2048"
  },
  "NetworkConfiguration": {
    "EgressConfiguration": {
      "EgressType": "",
      "VpcConnectorArn": null
    },
    "IngressConfiguration": {
      "IsPubliclyAccessible": true
    }
  },
  "ObservabilityConfiguration": {
    "ObservabilityEnabled": false,
    "ObservabilityConfigurationArn": null
  }
}
