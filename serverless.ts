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
      STAGE: '${opt:stage, "dev"}', // 環境変数
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
// aws ecr describe-images \
//     --repository-name sauna-api \
//     --profile private \
//     --image-ids imageDigest=sha256:81daaeb84084a4031b1325485fe912b47701fff5097ace274ea3a0fa9aa09b3f \
//     --region ap-northeast-1
