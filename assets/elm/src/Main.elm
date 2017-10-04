module FamilyTree exposing (..)

import Types exposing (Person, Family, examplePerson, exampleFamily)

import Html.App exposing (beginnerProgram)
import Hmtl exposing (Html, text, div)
import Html.Attributes exposing (class, src)
import List exposing (map, repeat)

type Msg = Nothing

type alias Model = 
  { people : List Person
  , families : List Family
  }

detailView : Person -> Html Msg
detailView person =
  div [ class "person" ]
      [ text [ "hello world" ] [] ]

treeView : List Person -> Html Msg
treeView people =
  div [ class "tree" ]
      (map detailView people)

model = 
  { people = repeat 12 examplePerson 
  , families = repeat 10 exampleFamily
  }

view : Model -> Html Msg
view model =
  treeView model.people

main =
  beginnerProgram { model = model
                  , view = view
                  , update = update
                  }
