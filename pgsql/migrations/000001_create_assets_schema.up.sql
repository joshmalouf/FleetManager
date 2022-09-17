CREATE SCHEMA "assets";

CREATE TYPE "assets"."op_status" AS ENUM (
  'active',
  'inactive',
  'disposed'
);

CREATE TYPE "assets"."stages" AS ENUM (
  '1',
  '1-2',
  '2',
  '2-3',
  '3',
  '4'
);

CREATE TABLE "assets"."cmppkgs" (
  "id" bigserial PRIMARY KEY,
  "unit_number" varchar(6) UNIQUE NOT NULL,
  "stages" varchar NOT NULL DEFAULT 'unknown',
  "op_status" varchar NOT NULL DEFAULT 'active',
  "com_status" varchar NOT NULL DEFAULT 'available',
  "current_location" varchar NOT NULL DEFAULT 'unknown',
  "driver_id" int,
  "compressor_id" int,
  "cooler_id" int,
  "vessel_id" int,
  "drawing_ref" path,
  "bom" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assets"."drivers" (
  "id" bigserial PRIMARY KEY,
  "engine_id" int,
  "motor_id" int,
  "unit_id" int DEFAULT null,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assets"."engines" (
  "id" bigserial PRIMARY KEY,
  "make" varchar NOT NULL,
  "model" varchar NOT NULL,
  "serial_number" varchar UNIQUE NOT NULL,
  "unit_id" int DEFAULT null,
  "op_status" assets.op_status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assets"."motors" (
  "id" bigserial PRIMARY KEY,
  "make" varchar NOT NULL,
  "model" varchar NOT NULL,
  "serial_number" varchar UNIQUE NOT NULL,
  "unit_id" int DEFAULT null,
  "op_status" assets.op_status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assets"."compressors" (
  "id" bigserial PRIMARY KEY,
  "make" varchar NOT NULL,
  "model" varchar NOT NULL,
  "serial_number" varchar NOT NULL,
  "throws" int NOT NULL,
  "op_status" assets.op_status NOT NULL,
  "unit_id" int DEFAULT null,
  "cylinder_id" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assets"."cylinders" (
  "id" bigserial PRIMARY KEY,
  "make" varchar NOT NULL,
  "model" varchar NOT NULL,
  "bore" float NOT NULL,
  "mawp" int NOT NULL,
  "serial_number" varchar UNIQUE NOT NULL,
  "comp_id" int DEFAULT null,
  "op_status" assets.op_status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assets"."coolers" (
  "id" bigserial PRIMARY KEY,
  "make" varchar NOT NULL,
  "model" varchar NOT NULL,
  "size" int NOT NULL,
  "job_number" varchar UNIQUE NOT NULL,
  "op_status" assets.op_status NOT NULL,
  "unit_id" int,
  "section_id" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assets"."clr_sections" (
  "id" bigserial PRIMARY KEY,
  "make" varchar NOT NULL,
  "serial_number" varchar UNIQUE NOT NULL,
  "mawp" int NOT NULL,
  "num_tubes" int,
  "num_rows" int,
  "passes" int,
  "cooler_id" int DEFAULT null,
  "op_status" assets.op_status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assets"."vessels" (
  "id" bigserial PRIMARY KEY,
  "scrubber_id" int,
  "pls_bottle_id" int,
  "unit_id" int
);

CREATE TABLE "assets"."scrubbers" (
  "id" bigserial PRIMARY KEY,
  "serial_number" varchar UNIQUE NOT NULL,
  "mawp" int NOT NULL,
  "diameter" int,
  "length" int,
  "drawing_ref" path,
  "op_status" assets.op_status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assets"."pls_bottles" (
  "id" bigserial PRIMARY KEY,
  "serial_number" varchar UNIQUE NOT NULL,
  "mawp" int NOT NULL,
  "diameter" int,
  "length" int,
  "drawing_ref" path,
  "op_status" assets.op_status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now()),
  "unit_id" int
);

CREATE TABLE "assets"."clrpkgs" (
  "id" bigserial PRIMARY KEY,
  "unit_number" varchar(6) UNIQUE NOT NULL,
  "Size" varchar NOT NULL DEFAULT 'unknown',
  "op_status" varchar NOT NULL DEFAULT 'active',
  "com_status" varchar NOT NULL DEFAULT 'available',
  "current_location" varchar NOT NULL DEFAULT 'unknown',
  "driver_id" int,
  "cooler_id" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "assets"."cmppkgs" ("unit_number");

CREATE INDEX ON "assets"."engines" ("serial_number");

CREATE INDEX ON "assets"."motors" ("serial_number");

CREATE INDEX ON "assets"."motors" ("unit_id");

CREATE INDEX ON "assets"."compressors" ("serial_number");

CREATE INDEX ON "assets"."compressors" ("unit_id");

CREATE INDEX ON "assets"."cylinders" ("serial_number");

CREATE INDEX ON "assets"."cylinders" ("comp_id");

CREATE INDEX ON "assets"."coolers" ("job_number");

CREATE INDEX ON "assets"."clr_sections" ("serial_number");

CREATE INDEX ON "assets"."clr_sections" ("cooler_id");

CREATE INDEX ON "assets"."vessels" ("id");

CREATE INDEX ON "assets"."scrubbers" ("serial_number");

CREATE INDEX ON "assets"."pls_bottles" ("serial_number");

CREATE INDEX ON "assets"."clrpkgs" ("unit_number");

ALTER TABLE "assets"."cmppkgs" ADD FOREIGN KEY ("driver_id") REFERENCES "assets"."drivers" ("id");

ALTER TABLE "assets"."cmppkgs" ADD FOREIGN KEY ("compressor_id") REFERENCES "assets"."compressors" ("id");

ALTER TABLE "assets"."cmppkgs" ADD FOREIGN KEY ("cooler_id") REFERENCES "assets"."coolers" ("id");

ALTER TABLE "assets"."cmppkgs" ADD FOREIGN KEY ("vessel_id") REFERENCES "assets"."vessels" ("id");

ALTER TABLE "assets"."drivers" ADD FOREIGN KEY ("engine_id") REFERENCES "assets"."engines" ("id");

ALTER TABLE "assets"."drivers" ADD FOREIGN KEY ("motor_id") REFERENCES "assets"."motors" ("id");

ALTER TABLE "assets"."drivers" ADD FOREIGN KEY ("unit_id") REFERENCES "assets"."cmppkgs" ("id");

ALTER TABLE "assets"."engines" ADD FOREIGN KEY ("unit_id") REFERENCES "assets"."cmppkgs" ("id");

ALTER TABLE "assets"."motors" ADD FOREIGN KEY ("unit_id") REFERENCES "assets"."cmppkgs" ("id");

ALTER TABLE "assets"."compressors" ADD FOREIGN KEY ("unit_id") REFERENCES "assets"."compressors" ("id");

ALTER TABLE "assets"."compressors" ADD FOREIGN KEY ("cylinder_id") REFERENCES "assets"."cylinders" ("id");

ALTER TABLE "assets"."cylinders" ADD FOREIGN KEY ("comp_id") REFERENCES "assets"."compressors" ("id");

ALTER TABLE "assets"."coolers" ADD FOREIGN KEY ("unit_id") REFERENCES "assets"."cmppkgs" ("id");

ALTER TABLE "assets"."coolers" ADD FOREIGN KEY ("section_id") REFERENCES "assets"."clr_sections" ("id");

ALTER TABLE "assets"."clr_sections" ADD FOREIGN KEY ("cooler_id") REFERENCES "assets"."coolers" ("id");

ALTER TABLE "assets"."vessels" ADD FOREIGN KEY ("scrubber_id") REFERENCES "assets"."scrubbers" ("id");

ALTER TABLE "assets"."vessels" ADD FOREIGN KEY ("pls_bottle_id") REFERENCES "assets"."pls_bottles" ("id");

ALTER TABLE "assets"."cmppkgs" ADD FOREIGN KEY ("id") REFERENCES "assets"."vessels" ("unit_id");

ALTER TABLE "assets"."clrpkgs" ADD FOREIGN KEY ("driver_id") REFERENCES "assets"."drivers" ("id");

ALTER TABLE "assets"."clrpkgs" ADD FOREIGN KEY ("cooler_id") REFERENCES "assets"."coolers" ("id");
