# TYPES
`No required`

- Can nest types
## Scalar type
Number, boolean, string,bytes

### Boolean
- default: false

### Number
- default: 0

### String
- UTF8 or 7bit ASCII format

### Enum
- must start by the tag 0

## Default value
- number: 0
- boolean: false
- string: empty string
- byte: empty bytes
- enum: first value
- repeated: empty list

# Packages

`[packageName].Date`


# Protocol buffer data evolution

- can add new property and remove
- Forward compatible change: new .proto => read Old.proto
- Backward compatible change: old .proto => read new .proto


## Updating Protocol Rules
1. Don't change the numeric tags for any existing fields
2. You can add new fileds, and old code will just ignore them
3. Likewise, if the old / new code read unknow data, the default will take place
4. Fields can be removed, as long as the tag number is not used again in your updated message type. You may want to rename the field instead, perhaps adding the prefiex "OBSOLETE_", or make the tag reserved, so that future users of your .proto can't acciden'tally reuse the number
5. For data type changes (int32 to int64 for example, please refer to the documentaion) or add new field

### Adding fields
1. Let's add field to our schema
   - old
    ```
    message MyMessage {
      int32 id = 1;
    }
    ```
   - new
    ```
    message MyMessage {
      int32 id = 1;
      string firstName = 2;
    }
    ```
2. if that field is sent to old code, the old code will not know what that tag number corresponds to and the field will be ignored (or dropped)
3. Oppositely, if we read old data with the new code, the new field will not be found and the default value will be assumed (empty string)
4. **Default values should always be interpreted with care**

### Renaming Fields
1. Let's rename a field in our schema
   - old
    ```
    message MyMessage {
      int32 id = 1;
      string firstName = 2;
    }
    ```
   - new
    ```
    message MyMessage {
      int32 id = 1;
      string person_firstName = 2;
    }
    ```
2. In this case nothing changes1 Field names ca be changed freely
3. **Only the tag number is important for protobuf**

### Removing Fields
1. Let's remove a field in our schema
   - old
    ```
    message MyMessage {
      int32 id = 1;
      string firstName = 2;
    }
    ```
   - new
    ```
    message MyMessage {
      int32 id = 1;
    }
    ```
2. If old code doesn't find the field anymore, the default value will be used
3. Oppositely, if we read old data with new code, the deleted field will just be dropped
4. **Default values should always be interpreted with care**
#### Extremly careful
5. When removing a field, yu should **ALWAYS** reserve the tag and the name
   - old
    ```
    message MyMessage {
      int32 id = 1;
      string firstName = 2;
    }
    ```
   - new
    ```
    message MyMessage {
      reserved 2;
      int32 id = 1;
    }
    ```
6. This prevents the tag to be re-used and this prevents the name to be re-used
7. This is neccesary to prevent conflicts in the codebase
#### Removing Fields Make some fields obsolete
8. The alternative is that instead of removing a field, you rename it to `OBSOLETE_field_name`
9. The downside is that you may have to populate that field while your client get upgraded to use the newer field that replaces it (which has a new tag)

### Reserved Keywords
1. you can reserve `TAGS` and `FIELD NAMES`
2. You can't mix TAGS and FIELD NAMES in the same reserved statement
```
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}
```
3. We resere TAGS to prevent new fields from re-using tags(which would break old code at runtime)
4. We reserve FIELD NAMES to prevent code bugs
5. Do not ever remove any reserved tags

### Beware of Defaults

1. Defaults are great, but the are tricky to deal with
2. Defaults allow us to easily evolve protobuf files without breaking any existing or new codes
3. They also ensure we know that a field will always have a non-null values
4. But they are dangerous, beacues
   - **You cannot differentiate from a missing field or if a value equal to the default was set**
5. Make sure the default value doesn't have meaning for your business
6. Deal with default values in your code if needed (with if statments)

### Evolving Enumerations
1. Enumerations can evolve:
  - can add
  - can remove
  - can reserve
2. If the code doesn't know what the received Enum value corresponds to, the default value will be used
3. Therefore, I recommend you make the first value "UNKNOWN = 0"
