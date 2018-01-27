# diff

Are there benefits to being able to diff two MARC records? Possible use
cases might be:

 * Updating databases (compare old to new to generate appropriate insert,
    update, and/or delete statements).
 * Merging records (determine whether or not two records are equivalent).
    This may also require normalizing the data before comparing. Which
    fields to consider? Which fields to ignore?
