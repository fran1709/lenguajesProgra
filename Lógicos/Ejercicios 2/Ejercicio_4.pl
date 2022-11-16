conectado(i,a,1).
conectado(i,b,1).
conectado(a,i,1).
conectado(a,c,1).
conectado(a,d,1).
conectado(b,i,1).
conectado(b,c,1).
conectado(b,d,1).
conectado(c,a,1).
conectado(c,x,1).
conectado(c,d,1).
conectado(d,a,1).
conectado(d,a,1).
conectado(d,f,1).
conectado(d,b,1).
conectado(x,c,1).

%        a ---- c ---- x
%      /   \  /
%     /     \/
%   i       X
%     \    / \
%      \  /   \
%        b ---- d ---- f

ruta(Inicio,Fin,Ruta,Peso):- ruta_aux(Inicio,Fin,[],Ruta,Peso).
rutaAux(Fin,Fin,_,[Fin],0).
rutaAux(Inicio,Fin,Visitados,[Inicio|Resto],Peso) :-
    conectado(Inicio,Alguien,PesoAux),
    not(member(Alguien,Visitados)),
    rutaAux(Alguien,Fin,[Inicio|Visitados],Resto,PesoAux2),
    Peso is PesoAux + PesoAux2.

listaRutas(Inicio,Fin, Rutas):- 
	findall([Ruta,Peso],
	ruta(Inicio,Fin,Ruta,Peso),Rutas).

