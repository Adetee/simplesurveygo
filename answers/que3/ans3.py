from jsonschema import validate
schema = {
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

validate(http://json-schema.org/
  		{
		  "colors": [
		    {
		      "color": "black",
		      "category": "hue",
		      "type": "primary",
		      "code": {
		        "rgba": [255,255,255,1],
		        "hex": "#000"
		      }
}, schema)
