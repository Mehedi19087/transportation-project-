drop index if exists idx_products_company_id;

alter table products 
      drop column if exists company_id;