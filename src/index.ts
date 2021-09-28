import { secrets } from "./config"

import axios from "axios"

axios.defaults.baseURL = `${ secrets.https ? "https" : "http" }://${ secrets.host }:${ secrets.port }`

interface RemoteData {
  identity: string
  rawCommand: string
}

console.log("process.argv :", process.argv)
const rawCommand = process.argv.slice(2).join(" ")

console.log("rawCommand :", rawCommand)

const data: RemoteData = {
  rawCommand,
  identity: secrets.identity
}

axios.post("/api/command", data, {
  headers: {
    Authorization: `Bearer ${ secrets.token }`
  }
})
