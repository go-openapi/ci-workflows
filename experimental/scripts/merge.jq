reduce (.[] | to_entries | .[]) as {$key, $value} (
  {} ;
  .[$key] += $value
) | .[] |= unique
