module Types exposing (..)

type alias Person =
  { id : String
  , name : String
  , sex : String
  , events : List Event
  , attributes : List Attribute
  , mother : Person
  , father : Person
  , children : List Person
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
  { id = "I1"
  , name = "John Smith"
  , sex = "m"
  , events : []
  , attributes : []
  , mother : Nothing
  , father : Nothing
  , children : []
  }

exampleWoman =
  { id = "I2"
  , name = "Jane Smith"
  , sex = "f"
  , events : []
  , attributes : []
  , mother : Nothing
  , father : Nothing
  , children : []
  }

exampleFamily =
  { husband : exampleMan
  , wife : exampleWoman
  , children : []
  }
