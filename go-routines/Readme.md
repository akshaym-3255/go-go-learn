concurrency in golang

it is achived using go routines 
they have simple syntax 


[23-12 13:31] Le, Tam N
{
    "name": "dec22-api3",
    "description": "Create job via API",
    "job_type": "BACKUP",
    "is_active": true,
    "no_of_retention_copies": 1,
    "is_auto_abort_active": false,
    "dsa_job_definition" : {
        "job_objects": [
                {
                    "object_name": "antares_small_db",
                    "object_type": "DATABASE",
                    "parent_name": "DBC",
                    "parent_type": "DATABASE",
                    "include_all": false
                }
            ]
    }
}

