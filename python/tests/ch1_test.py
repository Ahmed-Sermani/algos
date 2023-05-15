import pytest
from ch1 import is_unique, check_permutations
from ch1.one_away import one_away
from ch1.palindrome import palindrome_permutation, palindrome_permutation_v2
from ch1.string_compression import string_compression
from ch1.urlify import urlify

@pytest.mark.parametrize(
    "s,expect",
    [
        ["ahmad", False],
        ["sermani", True],
        ["aaaaaaaaaaaaaa", False],
    ]
)
def test_is_unique(s, expect):
    assert is_unique.is_unique_with_hashmap(s) is expect
    assert is_unique.is_unique_with_bitvector(s) is expect
    assert is_unique.is_unique_without_additional_datastructure(s) is expect

@pytest.mark.parametrize(
    "s1,s2,expect",
    [
        ["ahmad", "ahamd", True],
        ["sermani", "sermano", False],
        ["aaaaaaaaaaaaaa", "aaaaaaaaaaaaaa", True],
        ["abc", "bca", True],
        ["aaa", "aa", False],
    ]
)
def test_check_permutation(s1, s2, expect):
    assert check_permutations.check_permutation_brute_force(s1, s2) is expect
    assert check_permutations.check_permutation_hashmap(s1, s2) is expect
    assert check_permutations.check_permutation_vector(s1, s2) is expect
    assert check_permutations.check_permutation_bitvector(s1, s2) is expect
    assert check_permutations.check_permutation_sort(s1, s2) is expect


@pytest.mark.parametrize(
    "s,expect",
    [
        ["ahm ad  ", "ahm%20ad"],
        ["ser mani  ", "ser%20mani"],
    ]
)
def test_urlify(s, expect):
    assert urlify(s) == expect


@pytest.mark.parametrize(
    "s,expect",
    [
        ["Tact Coa", True],
        ["ahmad", False],
        ["tactcoapapa", True]
    ]
)
def test_palindrome_permutation(s, expect):
    assert palindrome_permutation(s) is expect
    assert palindrome_permutation_v2(s) is expect

@pytest.mark.parametrize(
    ["s1", "s2", "expect"],
    [
        ["ahmad", "ahmd", True],
        ["pale", "ple", True],
        ["pales", "pale", True],
        ["pale", "bale", True],
        ["pale", "bae", False]
    ]
)
def test_one_away(s1, s2, expect):
    assert one_away(s1, s2) is expect

@pytest.mark.parametrize(
    "s,expect",
    [
        ["aabcccccaaa", "a2b1c5a3"],
        ["ahmad", "ahmad"],
    ]
)
def test_string_compression(s, expect):
    assert string_compression(s) == expect
