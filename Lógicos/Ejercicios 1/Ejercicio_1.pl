%Eliminando un elemento de una lista.

delete([],_,[]).
delete([A|T],A,T):-!.
delete([H|T],E,L):- 
	delete(T,A,S), 
	append([H],S,L).