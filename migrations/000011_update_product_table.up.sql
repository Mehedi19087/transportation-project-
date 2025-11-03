alter table products 
      add column company_id bigint;

create index idx_products_company_id on products(company_id);