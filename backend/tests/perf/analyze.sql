EXPLAIN ANALYZE
SELECT username, first_name, last_name, birthday, gender, interests, city
FROM users
WHERE LOWER(first_name) LIKE LOWER('cd%')
  AND LOWER(last_name) LIKE LOWER('cd%')
ORDER BY id;

"Sort  (cost=253.65..253.71 rows=25 width=74) (actual time=0.273..0.274 rows=4 loops=1)"
"  Sort Key: id"
"  Sort Method: quicksort  Memory: 25kB"
"  ->  Bitmap Heap Scan on users  (cost=155.43..253.07 rows=25 width=74) (actual time=0.250..0.261 rows=4 loops=1)"
"        Filter: ((lower(first_name) ~~ 'cd%'::text) AND (lower(last_name) ~~ 'cd%'::text))"
"        Heap Blocks: exact=4"
"        ->  Bitmap Index Scan on user_prefix_idx  (cost=0.00..155.43 rows=25 width=0) (actual time=0.223..0.223 rows=4 loops=1)"
"              Index Cond: ((lower(first_name) ~>=~ 'cd'::text) AND (lower(first_name) ~<~ 'ce'::text) AND (lower(last_name) ~>=~ 'cd'::text) AND (lower(last_name) ~<~ 'ce'::text))"
"Planning Time: 0.160 ms"
"Execution Time: 0.316 ms"
