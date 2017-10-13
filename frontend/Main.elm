module FamilyTree exposing (..)

import Types exposing (Person, Family, exampleMan, exampleWoman, exampleFamily)

import Html exposing (Html, text, p, h3, div, program)
import Html.Attributes exposing (class, src)
import List exposing (concat, map, repeat)
import Http

-- Model
type alias Model = 
  { people : List Person
  , families : List Family
  }

init : ( Model, Cmd Msg )
init = 
  (Model [] [], Fetch)

-- Messages
type Msg 
  = Fetch
  | NewFamilyList (Result Http.Error String)

-- View
detailView : Person -> Html Msg
detailView person =
  div [ class "person" ]
      [ h3 [] [ text ("Person " ++ person.id) ]
      , p []
        [ text (person.name ++ ", ")
        , text (person.sex ++ ", ")
        , text (person.mother ++ ", ")
        , text (person.father)
        ]
      ]

treeView : List Person -> Html Msg
treeView people =
  div [ class "tree" ]
      ( map detailView people )

view : Model -> Html Msg
view model =
  treeView model.people

-- Update
update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        NewFamilyList (Ok, listOfFamilyID) ->
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
