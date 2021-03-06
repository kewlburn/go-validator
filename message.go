package validator

// MessageMap is a map of string, that can be used as error message for ValidateStruct function.
var MessageMap = map[string]string{
	"accepted":           "The :attribute must be accepted.",
	"activeUrl":          "The :attribute is not a valid URL.",
	"after":              "The :attribute must be a date after :date.",
	"afterOrEqual":       "The :attribute must be a date after or equal to :date.",
	"alpha":              "The :attribute may only contain letters.",
	"alphaDash":          "The :attribute may only contain letters, numbers, dashes and underscores.",
	"alphaNum":           "The :attribute may only contain letters and numbers.",
	"array":              "The :attribute must be an array.",
	"before":             "The :attribute must be a date before :date.",
	"beforeOrEqual":      "The :attribute must be a date before or equal to :date.",
	"between.numeric":    "The :attribute must be between :min and :max.",
	"between.file":       "The :attribute must be between :min and :max kilobytes.",
	"between.string":     "The :attribute must be between :min and :max characters.",
	"between.array":      "The :attribute must have between :min and :max items.",
	"boolean":            "The :attribute field must be true or false.",
	"confirmed":          "The :attribute confirmation does not match.",
	"date":               "The :attribute is not a valid date.",
	"dateFormat":         "The :attribute does not match the format :format.",
	"different":          "The :attribute and :other must be different.",
	"digits":             "The :attribute must be :digits digits.",
	"digitsBetween":      "The :attribute must be between :min and :max digits.",
	"dimensions":         "The :attribute has invalid image dimensions.",
	"distinct":           "The :attribute field has a duplicate value.",
	"email":              "The :attribute must be a valid email address.",
	"exists":             "The selected :attribute is invalid.",
	"file":               "The :attribute must be a file.",
	"filled":             "The :attribute field must have a value.",
	"gt.numeric":         "The :attribute must be greater than :value.",
	"gt.file":            "The :attribute must be greater than :value kilobytes.",
	"gt.string":          "The :attribute must be greater than :value characters.",
	"gt.array":           "The :attribute must have greater than :value items.",
	"gte.numeric":        "The :attribute must be greater than or equal :value.",
	"gte.file":           "The :attribute must be greater than or equal :value kilobytes.",
	"gte.string":         "The :attribute must be greater than or equal :value characters.",
	"gte.array":          "The :attribute must have :value items or more.",
	"image":              "The :attribute must be an image.",
	"in":                 "The selected :attribute is invalid.",
	"inArray":            "The :attribute field does not exist in :other.",
	"integer":            "The :attribute must be an integer.",
	"ip":                 "The :attribute must be a valid IP address.",
	"ipv4":               "The :attribute must be a valid IPv4 address.",
	"ipv6":               "The :attribute must be a valid IPv6 address.",
	"json":               "The :attribute must be a valid JSON string.",
	"lt.numeric":         "The :attribute must be less than :value.",
	"lt.file":            "The :attribute must be less than :value kilobytes.",
	"lt.string":          "The :attribute must be less than :value characters.",
	"lt.array":           "The :attribute must have less than :value items.",
	"lte.numeric":        "The :attribute must be less than or equal :value.",
	"lte.file":           "The :attribute must be less than or equal :value kilobytes.",
	"lte.string":         "The :attribute must be less than or equal :value characters.",
	"lte.array":          "The :attribute must not have more than :value items.",
	"max.numeric":        "The :attribute may not be greater than :max.",
	"max.file":           "The :attribute may not be greater than :max kilobytes.",
	"max.string":         "The :attribute may not be greater than :max characters.",
	"max.array":          "The :attribute may not have more than :max items.",
	"mimes":              "The :attribute must be a file of type: :values.",
	"mimetypes":          "The :attribute must be a file of type: :values.",
	"min.numeric":        "The :attribute must be at least :min.",
	"min.file":           "The :attribute must be at least :min kilobytes.",
	"min.string":         "The :attribute must be at least :min characters.",
	"min.array":          "The :attribute must have at least :min items.",
	"notIn":              "The selected :attribute is invalid.",
	"notRegex":           "The :attribute format is invalid.",
	"numeric":            "The :attribute must be a number.",
	"present":            "The :attribute field must be present.",
	"regex":              "The :attribute format is invalid.",
	"required":           "The :attribute field is required.",
	"requiredIf":         "The :attribute field is required when :other is :value.",
	"requiredUnless":     "The :attribute field is required unless :other is in :values.",
	"requiredWith":       "The :attribute field is required when :values is present.",
	"requiredWithAll":    "The :attribute field is required when :values is present.",
	"requiredWithout":    "The :attribute field is required when :values is not present.",
	"requiredWithoutAll": "The :attribute field is required when none of :values are present.",
	"same":               "The :attribute and :other must match.",
	"size.numeric":       "The :attribute must be :size.",
	"size.file":          "The :attribute must be :size kilobytes.",
	"size.string":        "The :attribute must be :size characters.",
	"size.array":         "The :attribute must contain :size items.",
	"string":             "The :attribute must be a string.",
	"timezone":           "The :attribute must be a valid zone.",
	"unique":             "The :attribute has already been taken.",
	"uploaded":           "The :attribute failed to upload.",
	"url":                "The :attribute format is invalid.",
}
