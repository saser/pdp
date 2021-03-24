# Test cases for package `fieldmask`

*   [ ] `Update` with the following setup:
    ```
	# dst: empty message (note that `nested` is unset, which is different from being set to an empty message)
	{}
	# src: non-empty nested message
	{
	  nested: {
	    foo: "some value"
		bar: "some other value"
	  }
	}
	# mask: specifies a field within the nested message
	paths: "nested.foo"
	# want: dst should contain the nested message, but only one field of it
	{
	  nested: {
	    foo: "some value"
	  }
	}
	```
*   [ ] Error messages that contain the values for invalid fields
*   [ ] Full replacement using the `*` field mask
*   [ ] Error when `*` is used in any other way than being the single top-level path (i.e., doing `nested.*` is invalid)
*   [ ] Full replacment with nil/empty mask and all fields in `src` are set
*   [ ] Partial replacement due to field masks when all fields in `src` are set
*   [ ] Errors when unsupported paths are used
    *    [ ] Groups (I don't even know what they are, but I should find out so that I can detect them)
	*    [ ] Accessing map entries by their keys
	*    [ ] Wildcarding over repeated fields
