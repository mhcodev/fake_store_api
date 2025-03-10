------------------------------------------------------------------------------------
-- IMPORTANT !!!
------------------------------------------------------------------------------------
-- Ensure fake_Store_db database is created before running this script
-- After that, run this script

------------------------------------------------------------------------------------
-- USER TYPES TABLE
------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "tb_user_type"(
    "id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL
);

------------------------------------------------------------------------------------
-- USERS TABLE
------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "tb_users"(
    "id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "user_type_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "avatar" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "status" SMALLINT NOT NULL DEFAULT '1',
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

-- Add Foreign Key Constraint If Not Exists
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints 
        WHERE constraint_name = 'tb_users_user_type_id_foreign'
    ) THEN
        ALTER TABLE "tb_users" 
        ADD CONSTRAINT "tb_users_user_type_id_foreign" 
        FOREIGN KEY ("user_type_id") REFERENCES "tb_user_type"("id");
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Add Unique Constraint If Not Exists
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints 
        WHERE constraint_name = 'u_email_uniq'
    ) THEN
        ALTER TABLE "tb_users" 
        ADD CONSTRAINT "u_email_uniq" 
        UNIQUE (email);
    END IF;
END;
$$ LANGUAGE plpgsql;

------------------------------------------------------------------------------------
-- CATEGORIES TABLE
------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "tb_categories"(
    "id" SERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "status" SMALLINT NOT NULL DEFAULT '1'
);

-- Add PRIMARY KEY Constraint If Not Exists
DO $$
BEGIN
    -- Check if the primary key constraint already exists
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints 
        WHERE constraint_name = 'tb_categories_pkey'
    ) THEN
        -- Add the primary key constraint if it doesn't exist
        ALTER TABLE "tb_categories" 
        ADD CONSTRAINT "tb_categories_pkey" PRIMARY KEY ("id");
    END IF;
END;
$$ LANGUAGE plpgsql;


------------------------------------------------------------------------------------
-- PRODUCTS TABLE
------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "tb_products"(
    "id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "category_id" BIGINT NOT NULL,
    "sku" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "slug" VARCHAR(255) NOT NULL,
    "stock" INTEGER NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "price" DOUBLE PRECISION NOT NULL,
    "discount" DECIMAL(8, 2) NOT NULL,
    "status" SMALLINT NOT NULL DEFAULT '1',
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

DO $$
BEGIN
    -- Check if the foreign key constraint already exists
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.table_constraints tc
        JOIN information_schema.key_column_usage kcu 
            ON tc.constraint_name = kcu.constraint_name
        WHERE tc.constraint_type = 'FOREIGN KEY'
        AND tc.table_name = 'tb_products'
        AND tc.constraint_name = 'tb_products_category_id_foreign'
    ) THEN
        -- Add the foreign key constraint if it doesn't exist
        ALTER TABLE "tb_products" 
            ADD CONSTRAINT "tb_products_category_id_foreign" 
            FOREIGN KEY ("category_id") 
            REFERENCES "tb_categories"("id");
    END IF;
END;
$$ LANGUAGE plpgsql;

------------------------------------------------------------------------------------
-- PRODUCT IMAGES TABLE
------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "tb_product_images"(
    "id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "product_id" BIGINT NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "base_url" VARCHAR(255) NOT NULL,
    "status" SMALLINT NOT NULL DEFAULT '1',
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

DO $$
BEGIN
    -- Check if the foreign key constraint already exists
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.table_constraints tc
        JOIN information_schema.key_column_usage kcu 
            ON tc.constraint_name = kcu.constraint_name
        WHERE tc.constraint_type = 'FOREIGN KEY'
        AND tc.table_name = 'tb_product_images'
        AND tc.constraint_name = 'tb_product_images_id_foreign'
    ) THEN
        -- Add the foreign key constraint if it doesn't exist
        ALTER TABLE "tb_product_images" 
            ADD CONSTRAINT "tb_product_images_id_foreign" 
            FOREIGN KEY ("product_id") 
            REFERENCES "tb_products"("id");
    END IF;
END;
$$ LANGUAGE plpgsql;

------------------------------------------------------------------------------------
-- FILES TABLE
------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "tb_files"(
    "id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "original_name" VARCHAR(255) NOT NULL,
    "filename" VARCHAR(255) NOT NULL,
    "type" VARCHAR(50)NOT NULL,
    "url" VARCHAR(255) NOT NULL,
    "base_url" VARCHAR(255) NOT NULL,
    "status" SMALLINT NOT NULL DEFAULT '1',
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

------------------------------------------------------------------------------------
-- LOG TABLE
------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "tb_api_logs"(
    "id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "method" VARCHAR(255) NOT NULL,
    "version" VARCHAR(50)NOT NULL,
    "path" VARCHAR(255) NOT NULL,
    "status_code" INTEGER NOT NULL,
    "response_time" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL DEFAULT 0,
    "ip_address" VARCHAR(255) NOT NULL,
    "country" VARCHAR(255) NOT NULL,
    "status" SMALLINT NOT NULL DEFAULT '1',
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);