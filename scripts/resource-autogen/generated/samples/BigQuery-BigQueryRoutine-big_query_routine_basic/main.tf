/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

```hcl
resource "google_bigquery_dataset" "test" {
	dataset_id = "dataset_id"
}

resource "google_bigquery_routine" "sproc" {
  dataset_id = google_bigquery_dataset.test.dataset_id
  routine_id     = "routine_id"
  routine_type = "PROCEDURE"
  language = "SQL"
  definition_body = "CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);"
}
```
