module Fetch exposing (..)

import Json.Decode exposing (string, list, map, map2, map3, map8, field, Decoder)
import Types exposing (Person, Event, Attribute)
import Http exposing (get, send)

fetch =
  let 
      family = "api/family"
      individual = "api/individual"
  in 
      send NewFamilyList (get family (list string))
      send NewIndividualList (get individual (list string))

personDecoder : Decoder Person
personDecoder =
  map8 Person
    (field "id" string)
    (field "name" string)
    (field "sex" string)
    (field "events" (list eventDecoder))
    (field "attributes" (list attributeDecoder))
    (field "mother" string)
    (field "father" string)
    (field "children" (list string))


eventDecoder : Decoder Event
eventDecoder =
  map3 Event
    (field "evntType" string)
    (field "date" string)
    (field "place" string)

attributeDecoder : Decoder Attribute
attributeDecoder =
  map2 Attribute
    (field "tag" string)
    (field "value" string)
