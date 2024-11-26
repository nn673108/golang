SELECT tbl_products.product_id, tbl_products.product_name, tbl_categories.category_name
FROM tbl_products
INNER JOIN tbl_categories ON tbl_products.category_id = tbl_categories.category_id;