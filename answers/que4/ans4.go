package main

import (
  "encoding/json"
  "log"
  "os"
  "gopkg.in/mgo.v2"
  data "movies.json"
)


func main() {
  var v []interface{}
  if err := json.Unmarshal(data, &v); 
  if err != nil {
    panic(err)
  }
  c := mongoSess.DB("mydb").C("mycollection")
  if err := c.Insert(v...); 
  if err != nil {
    panic(err)
  }
}
