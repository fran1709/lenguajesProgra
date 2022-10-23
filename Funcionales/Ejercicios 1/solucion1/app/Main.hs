module Main (main) where

import Data.List 

filter :: (a -> Bool) -> [a] -> [a]
filter _ [] = []
filter p (x:xs)
    | p x       = x : filter p xs
    | otherwise = filter p xs

main :: IO ()
main = do
    let sub_cadenas = ["la casa", "el perro", "pintando la cerca"];
    let notNull x = not ("la") in filter notNull sub_cadenas
    print(notNull)