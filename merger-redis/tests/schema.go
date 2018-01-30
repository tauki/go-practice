package main

// schema generated using :
// https://swagger-toolbox.firebaseapp.com/

const Schema = `{
    "properties": {
        "type": {
            "type": "string"
        },
        "master": {
            "type": "number"
        },
        "entity": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "code": {
                        "type": "number"
                    },
                    "ref": {
                        "type": "string"
                    },
                    "display": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`
