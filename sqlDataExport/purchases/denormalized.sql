select
    Proveedor.nombre as nombre_proveedor,
    Proveedor.contacto as contacto_proveedor,
    Proveedor.direccion as direccion_proveedor,
    Producto.nombre as nombre_producto,
    Producto.descripcion as descripcion_producto,
    Producto.fecha_fabricacion as fecha_fabricacion_producto,
    Pais.nombre as nombre_paid,
    Categoria.nombre as nombre_categoria,
    Categoria.descripcion as descripcion_categoria,
    Subcategoria.nombre as nombre_subcategoria,
    Subcategoria.descripcion descripcion_subcategoria,
    Compra.fecha,
    Compra.cantidad_producto,
    Compra.precio_unitario,
    Compra.iva_pagado,
    Compra.subtotal,
    Compra.total
from compra as Compra
join proveedor Proveedor on Proveedor.id_proveedor = Compra.id_proveedor
join producto Producto on Producto.id_producto = Compra.id_producto
join pais Pais on Pais.id_pais = Compra.id_pais
join subcategoria Subcategoria on Subcategoria.id_subcategoria = Compra.id_subcategoria
join categoria Categoria on Categoria.id_categoria = Subcategoria.id_categoria