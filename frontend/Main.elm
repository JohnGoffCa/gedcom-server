module FamilyTree exposing (..)

import Types exposing (Person, Family, exampleMan, exampleFamily)

import Html exposing (Html, text, div, beginnerProgram)
import Html.Attributes exposing (class, src)
import List exposing (map, repeat)

type Msg 
  = NoOp

type alias Model = 
  { people : List Person
  , families : List Family
  }

detailView : Person -> Html Msg
detailView person =
  div [ class "person" ]
      [ text "hello world" ]

treeView : List Person -> Html Msg
treeView people =
  div [ class "tree" ]
      (map detailView people)

model = 
  { people = repeat 12 exampleMan 
  , families = repeat 10 exampleFamily
  }

view : Model -> Html Msg
view model =
  treeView model.people

update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        NoOp ->
            ( model, Cmd.none )

main =
  beginnerProgram { model : model
                  , view : view
                  , update : update
                  }
