from jsonschema import validate
{
        "$schema":"http://json-schema.org#",
        "type":"object",
        "properties": {
            "colors" : {
                "type":"array",
                "items" : {
                    "type":"object",
                    "properties": {
                        "color" :{"type":"string"},
                        "category" : {"type":"string"},
                        "type": {"type":"string"},
                        "code": {
                            "type":"array",
                            "items":{
                                "type":"object",
                                "properties":{
                                    "rgba":{"type":"array"},
                                    "hex":{"type":"string"}
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }


