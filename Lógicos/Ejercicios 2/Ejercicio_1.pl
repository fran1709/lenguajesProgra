subcadena(Elemento, Texto) :- sub_cadena(Texto, _, _, _, Valor), !.
filtrar_lista(Elemento, Aprove, Filter) :-
        include(subcadena(Element), Aprove, Filter).