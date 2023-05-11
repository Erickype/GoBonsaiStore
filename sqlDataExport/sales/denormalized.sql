select
    Cliente.nombre as nombre_cliente, apellido as apellido_cliente,
    Pais.nombre as nombre_pais,
    Subcategoria.nombre nombre_subcategoria, Subcategoria.descripcion as descripcion_subcategoria,
    Producto.nombre as nombre_producto, Producto.fecha_fabricacion, Producto.descripcion as descripcion_producto,
    Venta.fecha,cantidad_producto,precio_unitario,iva_cobrado,subtotal,total
from venta as Venta
         join cliente Cliente on Cliente.id_cliente = Venta.id_cliente
         join pais Pais on Pais.id_pais = Venta.id_pais
         join subcategoria Subcategoria on Subcategoria.id_subcategoria = Venta.id_subcategoria
         join producto Producto on Producto.id_producto = Venta.id_producto
