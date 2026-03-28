package stdlib

func RunAll() {
	println("=== String Operations ===")
	result := StringOperations()
	println(result)

	println("\n=== Sort Operations ===")
	nums := SortOperations()
	println("Sorted:", nums)

	println("\n=== Slice Operations ===")
	same := SliceOperations()
	println("Equal:", same)

	println("\n=== JSON Operations ===")
	name, _ := JSONOperations()
	println("User:", name)
	size, _ := JSONWithOmitEmpty()
	println("JSON size with omit:", size)

	println("\n=== File Operations ===")
	FileOperations()
	println("File operations done")

	println("\n=== Path Operations ===")
	path := PathOperations()
	println(path)

	println("\n=== Time Operations ===")
	timeResult := TimeOperations()
	println(timeResult)
	duration := DurationOperations()
	println(duration)

	println("\n=== Utils ===")
	strResult, _ := StrconvExamples()
	println(strResult)
	println("Bytes equal:", BytesOperations())
	println("Regex match:", RegexpExamples())
	println("Buffer:", BufferExample())
}
