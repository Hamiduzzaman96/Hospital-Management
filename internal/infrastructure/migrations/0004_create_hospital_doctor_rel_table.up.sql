CREATE TABLE hospital_doctor_rel (
    hospital_id BIGINT NOT NULL,
    doctor_id BIGINT NOT NULL,
    PRIMARY KEY (hospital_id, doctor_id),
    FOREIGN KEY (hospital_id) REFERENCES hospitals(id) ON DELETE CASCADE,
    FOREIGN KEY (doctor_id) REFERENCES doctors(id) ON DELETE CASCADE
);

CREATE INDEX idx_hospital_doctor_hospital_id ON hospital_doctor_rel(hospital_id);
CREATE INDEX idx_hospital_doctor_doctor_id ON hospital_doctor_rel(doctor_id);