// eslint-disable-next-line @typescript-eslint/no-var-requires
const secrets = require("../secrets.json") as Secrets

interface Secrets {
  host: string
  port: number | string
  identity: string
  token: string
  https: boolean
}

export {
  secrets
}
