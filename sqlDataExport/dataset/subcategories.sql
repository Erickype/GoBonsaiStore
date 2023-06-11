select
    C.edad, G.id as sexo, P.id as profesion, C.salario,
    V.cantidad_producto, V.precio_unitario, V.total,
    V.id_subcategoria as clase
from hechos.venta V
         join dimensiones.cliente C on V.id_cliente = C.id_cliente
         join views.gender G on C.sexo = G.sexo
         join views.profession P on C.profesion = P.profesion;