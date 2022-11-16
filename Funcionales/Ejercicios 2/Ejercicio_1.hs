data Producto = Producto {
    nombre :: String,
    info :: String,
    precio :: Int
} deriving Show

reducir :: (a -> a -> a) -> [a] -> a
reducir d [a] = a
reducir d xs = reduce d ((d (head xs) (xs !! 1):drop 2 xs))
                     
monto_factura_total lista_productos=
    map articulos lista_productos
    where
        productos producto = ((precio producto))

inicio :: IO ()
inicio = do
    let impuesto = 20000
    let porcentaje = 13
    
    let menu = [
            (Producto "pizza" "comida" 12500),
            (Producto "hamburguesa" "comida" 2300),
            (Producto "pollo frito" "comida" 1200),
            (Producto "papanachos" "comida" 2000),
            (Producto "Coca Cola" "bebida" 1000)]

print("Total a pagar sin I.V.A: ")
    print(reducir (+) (monto_factura_total menu))
print("Total a pagar con I.V.A:")
    print((reduce (+) (monto_factura_total menu)) + ((reducir (+) (map (porcentaje *) (filter (> impuesto) (monto_factura_total menu)))) `div` 100))