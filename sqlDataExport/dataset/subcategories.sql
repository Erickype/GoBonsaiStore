select
    V.id_pais as pais, C.edad, G.id as sexo, P.id as profesion, C.salario,
    SubCategoria.id_categoria as categoria,
    DATE_PART('year', CURRENT_DATE) - DATE_PART('year', Producto.fecha_fabricacion) as edad_producto,
    V.cantidad_producto, V.precio_unitario, V.total,
    V.id_subcategoria as clase
from hechos.venta V
join dimensiones.cliente C on V.id_cliente = C.id_cliente
join views.gender G on C.sexo = G.sexo
join views.profession P on C.profesion = P.profesion
join dimensiones.subcategoria SubCategoria on V.id_subcategoria = SubCategoria.id_subcategoria
join dimensiones.producto Producto on V.id_producto = Producto.id_producto
order by V.fecha;

