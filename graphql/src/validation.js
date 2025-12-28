import { buildSchema, validateSchema } from "graphql"
import fs from "fs"

const schemaSDL = fs.readFileSync("schema.graphql", "utf8")
const schema = buildSchema(schemaSDL)

const errors = validateSchema(schema)
console.log(errors)
