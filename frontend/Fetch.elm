module Fetch exposing (..)

import Json.Decode exposing (string, list, nullable, map7, at, Decoder)
import Json.Decode.Pipeline exposing (decode, requiredAt, required, optional, hardcoded)
import Types exposing (Person, Event, Attribute)

personDecoder : Decoder Person
personDecoder =
  decode Person
  |> required "id" string
  |> required "name" string
  |> required "sex" string
  |> requiredAt ["events"] (list eventDecoder)
  |> required "mother" (nullable string)
  |> required "father" (nullable string)
  |> required "children" (list string)


eventListDecoder : Decoder (List Event)
eventListDecoder =
  list eventDecoder

eventDecoder : Decoder Event
eventDecoder =
  decode Event
  |> requiredAt ["events", "evntType"] string
  |> requiredAt ["events", "date"] string
  |> requiredAt ["events", "place"] string
