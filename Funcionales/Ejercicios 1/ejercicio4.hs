data Abrol a = Hoja a | Nodo [Arbol a]

aplanar :: Arbol a -> [a]
aplanar (Hojax) = [x]
aplanar (Nodo xs) = xs >>= aplanar

main :: IO ()
main = do
    let Tree = Nodo [
		  Nodo [Hoja 1], 
		  Hoja 2, 
		  Nodo [
		    Nodo [Hoja 3, Hoja 4], 
		    Hoja 5
		  ], 
		  Nodo [Nodo [Nodo []]], 
		  Nodo [Nodo [Nodo [Hoja 6]]], 
		  Hoja 7, 
		  Hoja 8, 
		  Nodo []
                ]

    print(aplanar Tree)