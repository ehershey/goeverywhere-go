result = db.gps_log.ensureIndex({ entry_date: 1, entry_source: 1, accuracy: 1 }, { unique: true, dropDups: true });
printjson(result);
if(!result.ok) quit(1)
result = db.gps_log.ensureIndex( { loc : "2dsphere" } );
printjson(result);
if(!result.ok) quit(1)
result = db.gps_log.ensureIndex({ entry_source: 1 });
printjson(result);
if(!result.ok) quit(1)
result = db.nodes.ensureIndex(   {  loc: '2dsphere' } );
printjson(result);
if(!result.ok) quit(1)
result = db.nodes.ensureIndex( { external_id: 1 },  { unique: true });
printjson(result);
if(!result.ok) quit(1)
result = db.jobs.ensureIndex({ job_id: 1 }, { unique: true });
printjson(result);
if(!result.ok) quit(1)
