-- Create "admin_users" table
CREATE TABLE "public"."admin_users" (
  "id" bigserial NOT NULL,
  "email" character varying(255) NULL,
  "password_hash" character varying(255) NULL,
  "full_name" character varying(255) NULL,
  "role" character varying(50) NULL,
  "telegram_id" character varying(100) NOT NULL,
  "is_active" boolean NULL DEFAULT true,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_admin_users_email" to table: "admin_users"
CREATE UNIQUE INDEX "idx_admin_users_email" ON "public"."admin_users" ("email");
-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" bigserial NOT NULL,
  "name" character varying(100) NULL,
  "slug" character varying(100) NULL,
  "parent_id" bigint NULL,
  "sort_order" bigint NULL DEFAULT 0,
  "is_active" boolean NULL DEFAULT true,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_categories_children" FOREIGN KEY ("parent_id") REFERENCES "public"."categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_categories_slug" to table: "categories"
CREATE UNIQUE INDEX "idx_categories_slug" ON "public"."categories" ("slug");
-- Create "products" table
CREATE TABLE "public"."products" (
  "id" bigserial NOT NULL,
  "name" character varying(255) NULL,
  "slug" character varying(255) NULL,
  "description" text NULL,
  "price" numeric(10,2) NULL,
  "old_price" numeric(10,2) NULL,
  "articule" character varying(100) NULL,
  "category_id" bigint NULL,
  "is_active" boolean NULL DEFAULT true,
  "is_available" boolean NULL DEFAULT true,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_products_category" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_products_articule" to table: "products"
CREATE UNIQUE INDEX "idx_products_articule" ON "public"."products" ("articule");
-- Create index "idx_products_slug" to table: "products"
CREATE UNIQUE INDEX "idx_products_slug" ON "public"."products" ("slug");
-- Create "product_sizes" table
CREATE TABLE "public"."product_sizes" (
  "id" bigserial NOT NULL,
  "product_id" bigint NULL,
  "size_name" character varying(50) NULL,
  "quantity" bigint NULL DEFAULT 0,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_products_sizes" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "email" character varying(255) NULL,
  "phone" character varying(20) NULL,
  "password_hash" character varying(255) NULL,
  "full_name" character varying(255) NULL,
  "telegram_id" character varying(100) NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create index "idx_users_phone" to table: "users"
CREATE UNIQUE INDEX "idx_users_phone" ON "public"."users" ("phone");
-- Create "cart_items" table
CREATE TABLE "public"."cart_items" (
  "id" bigserial NOT NULL,
  "user_id" bigint NULL,
  "product_id" bigint NULL,
  "product_size_id" bigint NULL,
  "quantity" bigint NULL DEFAULT 1,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_cart_items_product" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_cart_items_product_size" FOREIGN KEY ("product_size_id") REFERENCES "public"."product_sizes" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_cart_items_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "orders" table
CREATE TABLE "public"."orders" (
  "id" bigserial NOT NULL,
  "order_number" character varying(50) NULL,
  "user_id" bigint NULL,
  "status" character varying(50) NULL,
  "total_amount" numeric(10,2) NULL,
  "customer_name" character varying(255) NULL,
  "customer_phone" character varying(20) NULL,
  "customer_email" character varying(255) NULL,
  "customer_telegram" character varying(100) NULL,
  "pickup_point" character varying(500) NULL,
  "notes" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_orders_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_orders_order_number" to table: "orders"
CREATE UNIQUE INDEX "idx_orders_order_number" ON "public"."orders" ("order_number");
-- Create "order_items" table
CREATE TABLE "public"."order_items" (
  "id" bigserial NOT NULL,
  "order_id" bigint NULL,
  "product_id" bigint NULL,
  "product_size_id" bigint NULL,
  "product_name" character varying(255) NULL,
  "product_price" numeric(10,2) NULL,
  "quantity" bigint NULL,
  "total_price" numeric(10,2) NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_order_items_product" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_order_items_product_size" FOREIGN KEY ("product_size_id") REFERENCES "public"."product_sizes" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_orders_items" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "order_messages" table
CREATE TABLE "public"."order_messages" (
  "id" bigserial NOT NULL,
  "order_id" bigint NULL,
  "sender_type" character varying(20) NULL,
  "message" text NULL,
  "is_read" boolean NULL DEFAULT false,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_orders_messages" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "product_images" table
CREATE TABLE "public"."product_images" (
  "id" bigserial NOT NULL,
  "product_id" bigint NULL,
  "image_path" character varying(500) NULL,
  "sort_order" bigint NULL DEFAULT 0,
  "is_main" boolean NULL DEFAULT false,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_products_images" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "wishlist" table
CREATE TABLE "public"."wishlist" (
  "id" bigserial NOT NULL,
  "user_id" bigint NULL,
  "product_id" bigint NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_wishlist_product" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_wishlist_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
