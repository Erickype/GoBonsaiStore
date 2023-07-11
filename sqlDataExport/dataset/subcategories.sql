select
    V.id_pais as pais, C.edad, G.id as sexo, P.id as profesion, C.salario,
    SubCategoria.id_subcategoria as subcategoria,
    DATE_PART('year', CURRENT_DATE) - DATE_PART('year', Producto.fecha_fabricacion) as edad_producto,
    Producto.alto, Producto.ancho, Producto.profundidad,
    CASE
        WHEN Producto.edad_relevante = TRUE THEN 1
        ELSE 0
        END AS edad_relevante,
    V.cantidad_producto, V.precio_unitario, V.iva_cobrado as iva, V.subtotal, V.total, Ganancias.ganancia,
    SubCategoria.id_categoria as clase
from hechos.venta V
         join dimensiones.cliente C on V.id_cliente = C.id_cliente
         join views.gender G on C.sexo = G.sexo
         join views.profession P on C.profesion = P.profesion
         join dimensiones.subcategoria SubCategoria on V.id_subcategoria = SubCategoria.id_subcategoria
         join dimensiones.producto Producto on V.id_producto = Producto.id_producto
         join hechos.ganancias Ganancias on
            V.id_producto = Ganancias.id_producto
        and V.id_cliente = Ganancias.id_cliente
        and V.fecha = Ganancias."fechaVenta"
order by V.fecha;