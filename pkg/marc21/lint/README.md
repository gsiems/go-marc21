# lint

It is desired to be able to read a MARC record and identify the following
"bad data".
 * Leader values that are incorrect (not in appropriate list)
 * Controlfield values that are incorrect (not in appropriate list)
 * Required fields that are missing
 * Non-repeating fields that repeat
 * Non-repeating subfields that repeat
 * Incorrect indicators for datafields
 * Datafields that have invalid subfields
 * Obsolete tags

This linting needs to be context aware regarding the MARC record format.
That is, the rules need to be adjusted depending on whether the record is
for [Authority](http://www.loc.gov/marc/authority/),
[Bibliography](http://www.loc.gov/marc/bibliographic/),
[Classification](http://www.loc.gov/marc/classification/),
[Community Information](http://www.loc.gov/marc/community/), or
[Holdings](http://www.loc.gov/marc/holdings/) records.

Is it practical to attempt correcting some forms of bad data? Adding
missing data is out, and removing empty fields **should** be a
no-brainer (however this is MARC after all, so... I'm probably not
understanding something), but is it practical to scrub the
duplicated/empty data? Remove the repeated entries of non-repeating
fields/subfields (which one to keep? the first, last? how do various
other tools deal with this? is there a standard convention in the face
of repeated non-repeating elements?). Can some obsolete fields be
migrated to the appropriate replacement fields?

How easily can (some of) the rules be generated from the following:
 * http://www.loc.gov/marc/bibliographic/ecbdlist.html
 * http://www.loc.gov/marc/holdings/echdlist.html
 * http://www.loc.gov/marc/authority/ecadlist.html
 * http://www.loc.gov/marc/classification/eccdlist.html
 * http://www.loc.gov/marc/community/eccilist.html
