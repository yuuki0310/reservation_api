import type { AWS } from "@serverless/typescript";

const serverlessConfiguration: AWS = {
  service: "sauna-api", // サービス名
  frameworkVersion: "3",
  provider: {
    name: "aws",
    runtime: "provided.al2", // Amazon Linux 2を使用
    region: "ap-northeast-1", // リージョン
    architecture: "arm64", // Lambda関数のアーキテクチャ
    environment: {
      ENV: '${opt:dev, "dev"}', // 環境変数
    },
    vpc: {
      securityGroupIds: [
        "sg-0962d6ce4b1084ea8", // Lambda用のセキュリティグループID
      ],
      subnetIds: [
        "subnet-001bc4194590626c5",
        "subnet-08b270ac475e57943",
        "subnet-0f42728727be4389f",
      ],
    },
    httpApi: {
      cors: {
        allowedOrigins: ["*"], // すべてのオリジンを許可
        allowedHeaders: ["Content-Type", "Authorization"], // 許可するヘッダー
        allowedMethods: ["GET", "POST", "OPTIONS"], // 許可するメソッド
      },
    },
  },
  custom: {
    region: "${opt:region, 'ap-northeast-1'}",
  },
  functions: {
    api: {
      image: {
        uri: "${aws:accountId}.dkr.ecr.${self:custom.region}.amazonaws.com/sauna-api:latest",
      },
      events: [
        {
          httpApi: {
            path: "/",
            method: "get",
          },
        },
        {
          httpApi: {
            path: "/stores/{storeId}/reservations",
            method: "get",
          },
        },
        {
          httpApi: {
            path: "/reservations",
            method: "post",
          },
        },
        {
          httpApi: {
            path: "/users/{uuid}/reservations",
            method: "get",
          },
        },
      ],
    },
  },
  plugins: ["serverless-offline"],
};

module.exports = serverlessConfiguration;
