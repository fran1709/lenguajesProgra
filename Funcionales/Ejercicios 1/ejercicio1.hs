module ExerciseOne(main) where
    
import Data.List 
import Data.Char

filtrarSubcadena:: String -> [String] -> [String]
filtrarSubcadena = filter xs

main::IO()
main = do
    let sub_cadenas = ["la casa", "el perro", "pintando la cerca"]
    filtrarSubcaden "la" sub_cadenas
