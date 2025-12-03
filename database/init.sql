CREATE TABLE batch_inputs (
    sl_no INT AUTO_INCREMENT PRIMARY KEY,
    start_time DATETIME,
    raw_material_name VARCHAR(50),
    raw_material_quantity INT,
    temperature INT,
    end_time DATETIME
)

SELECT 'Database initialized successfully!' AS status;