module Types exposing (..)

type alias Person =
  { id : String
  , name : String
  , sex : String
  , events : List Event
  , attributes : List Attribute
  , mother : String
  , father : String
  , children : List String
  }

type alias Event =
  { evntType : String
  , date : String
  , place : String
  }

type alias Attribute =
  { attrType : String
  , value : String
  }

type alias Family =
  { husband : Person
  , wife : Person
  , children : List Person
  }

exampleMan =
  Person "I1" "John Smith" "m" [] [] "I10" "I11" []

exampleWoman =
  Person "I2" "Jane Smith" "f" [] [] "I10" "I11" []

exampleFamily =
  Family exampleMan exampleWoman []
