module FamilyTree exposing (..)

import Types exposing (Person, Family, exampleMan, exampleFamily)

import Html exposing (Html, text, div, program)
import Html.Attributes exposing (class, src)
import List exposing (map, repeat)

-- Model
type alias Model = 
  { people : List Person
  , families : List Family
  }

init : ( Model, Cmd Msg )
init = 
  (Model (repeat 12 exampleMan) (repeat 10 exampleFamily), Cmd.none)

-- Messages
type Msg 
  = NoOp

-- View
detailView : Person -> Html Msg
detailView person =
  div [ class "person" ]
      [ text "hello world" ]

treeView : List Person -> Html Msg
treeView people =
  div [ class "tree" ]
      (map detailView people)

view : Model -> Html Msg
view model =
  treeView model.people

-- Update
update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        NoOp ->
            ( model, Cmd.none )

--subscriptions 
subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none

--main
main : Program Never Model Msg
main =
  program
    { init = init
    , view = view
    , update = update
    , subscriptions = subscriptions
    }
