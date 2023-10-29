// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/avstring.h>
*/
import "C"

// NONEED: av_strstart

// NONEED: av_stristart

// NONEED: av_stristr

// NONEED: av_strnstr

// NONEED: av_strlcpy

// NONEED: av_strlcat

// NONEED: av_strlcatf

// NONEED: av_strnlen

// NONEED: av_asprintf

// NONEED: av_d2str

// NONEED: av_get_token

// NONEED: av_strtok

// NONEED: av_isgraph

// NONEED: av_isspace

// NONEED: av_toupper

// NONEED: av_tolower

// NONEED: av_isxdigit

// NONEED: av_strcasecmp

// NONEED: av_strncasecmp

// NONEED: av_strireplace

// NONEED: av_basename

// NONEED: av_dirname

// NONEED: av_match_name

// NONEED: av_append_path_component

type AVEscapeMode = C.enum_AVEscapeMode

const (
	AV_ESCAPE_MODE_AUTO      = C.AV_ESCAPE_MODE_AUTO
	AV_ESCAPE_MODE_BACKSLASH = C.AV_ESCAPE_MODE_BACKSLASH
	AV_ESCAPE_MODE_QUOTE     = C.AV_ESCAPE_MODE_QUOTE
	AV_ESCAPE_MODE_XML       = C.AV_ESCAPE_MODE_XML
)

const (
	AV_ESCAPE_FLAG_WHITESPACE        = C.AV_ESCAPE_FLAG_WHITESPACE
	AV_ESCAPE_FLAG_STRICT            = C.AV_ESCAPE_FLAG_STRICT
	AV_ESCAPE_FLAG_XML_SINGLE_QUOTES = C.AV_ESCAPE_FLAG_XML_SINGLE_QUOTES
	AV_ESCAPE_FLAG_XML_DOUBLE_QUOTES = C.AV_ESCAPE_FLAG_XML_DOUBLE_QUOTES
)

// NONEED: av_escape

const (
	AV_UTF8_FLAG_ACCEPT_INVALID_BIG_CODES          = C.AV_UTF8_FLAG_ACCEPT_INVALID_BIG_CODES
	AV_UTF8_FLAG_ACCEPT_NON_CHARACTERS             = C.AV_UTF8_FLAG_ACCEPT_NON_CHARACTERS
	AV_UTF8_FLAG_ACCEPT_SURROGATES                 = C.AV_UTF8_FLAG_ACCEPT_SURROGATES
	AV_UTF8_FLAG_EXCLUDE_XML_INVALID_CONTROL_CODES = C.AV_UTF8_FLAG_EXCLUDE_XML_INVALID_CONTROL_CODES
	AV_UTF8_FLAG_ACCEPT_ALL                        = C.AV_UTF8_FLAG_ACCEPT_ALL
)

// NONEED: av_utf8_decode

// NONEED: av_match_list

// NONEED: av_sscanf
