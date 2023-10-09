package hashmap

import "github.com/mothdb-bd/common-go/pkg/basic"

// Java Map 接口兼容
type HashMap[K any, V any] interface {
	// Query Operations

	/**
	* Returns the number of key-value mappings in this map.  If the
	* map contains more than {@code Integer.MAX_VALUE} elements, returns
	* {@code Integer.MAX_VALUE}.
	*
	* @return the number of key-value mappings in this map
	 */
	Size() int32

	/**
	* Returns {@code true} if this map contains no key-value mappings.
	*
	* @return {@code true} if this map contains no key-value mappings
	 */
	IsEmpty() bool

	/**
	* Returns {@code true} if this map contains a mapping for the specified
	* key.  More formally, returns {@code true} if and only if
	* this map contains a mapping for a key {@code k} such that
	* {@code Objects.equals(key, k)}.  (There can be
	* at most one such mapping.)
	*
	* @param key key whose presence in this map is to be tested
	* @return {@code true} if this map contains a mapping for the specified
	*         key
	* @throws ClassCastException if the key is of an inappropriate type for
	*         this map
	* (<a href="{@docRoot}/java.base/java/util/Collection.html#optional-restrictions">optional</a>)
	* @throws NullPointerException if the specified key is null and this map
	*         does not permit null keys
	* (<a href="{@docRoot}/java.base/java/util/Collection.html#optional-restrictions">optional</a>)
	 */
	ContainsKey(key basic.Object) bool

	/**
	* Returns {@code true} if this map maps one or more keys to the
	* specified value.  More formally, returns {@code true} if and only if
	* this map contains at least one mapping to a value {@code v} such that
	* {@code Objects.equals(value, v)}.  This operation
	* will probably require time linear in the map size for most
	* implementations of the {@code HashMap} interface.
	*
	* @param value value whose presence in this map is to be tested
	* @return {@code true} if this map maps one or more keys to the
	*         specified value
	* @throws ClassCastException if the value is of an inappropriate type for
	*         this map
	* (<a href="{@docRoot}/java.base/java/util/Collection.html#optional-restrictions">optional</a>)
	* @throws NullPointerException if the specified value is null and this
	*         map does not permit null values
	* (<a href="{@docRoot}/java.base/java/util/Collection.html#optional-restrictions">optional</a>)
	 */
	ContainsValue(value basic.Object) bool

	/**
	* Returns the value to which the specified key is mapped,
	* or {@code null} if this map contains no mapping for the key.
	*
	* <p>More formally, if this map contains a mapping from a key
	* {@code k} to a value {@code v} such that
	* {@code Objects.equals(key, k)},
	* then this method returns {@code v}; otherwise
	* it returns {@code null}.  (There can be at most one such mapping.)
	*
	* <p>If this map permits null values, then a return value of
	* {@code null} does not <i>necessarily</i> indicate that the map
	* contains no mapping for the key; it's also possible that the map
	* explicitly maps the key to {@code null}.  The {@link #containsKey
	* containsKey} operation may be used to distinguish these two cases.
	*
	* @param key the key whose associated value is to be returned
	* @return the value to which the specified key is mapped, or
	*         {@code null} if this map contains no mapping for the key
	* @throws ClassCastException if the key is of an inappropriate type for
	*         this map
	* (<a href="{@docRoot}/java.base/java/util/Collection.html#optional-restrictions">optional</a>)
	* @throws NullPointerException if the specified key is null and this map
	*         does not permit null keys
	* (<a href="{@docRoot}/java.base/java/util/Collection.html#optional-restrictions">optional</a>)
	 */
	Get(key basic.Object) V

	// Modification Operations

	/**
	* Associates the specified value with the specified key in this map
	* (optional operation).  If the map previously contained a mapping for
	* the key, the old value is replaced by the specified value.  (A map
	* {@code m} is said to contain a mapping for a key {@code k} if and only
	* if {@link #containsKey(Object) m.containsKey(k)} would return
	* {@code true}.)
	*
	* @param key key with which the specified value is to be associated
	* @param value value to be associated with the specified key
	* @return the previous value associated with {@code key}, or
	*         {@code null} if there was no mapping for {@code key}.
	*         (A {@code null} return can also indicate that the map
	*         previously associated {@code null} with {@code key},
	*         if the implementation supports {@code null} values.)
	* @throws UnsupportedOperationException if the {@code put} operation
	*         is not supported by this map
	* @throws ClassCastException if the class of the specified key or value
	*         prevents it from being stored in this map
	* @throws NullPointerException if the specified key or value is null
	*         and this map does not permit null keys or values
	* @throws IllegalArgumentException if some property of the specified key
	*         or value prevents it from being stored in this map
	 */
	Put(key K, value V) V

	/**
	* Removes the mapping for a key from this map if it is present
	* (optional operation).   More formally, if this map contains a mapping
	* from key {@code k} to value {@code v} such that
	* {@code Objects.equals(key, k)}, that mapping
	* is removed.  (The map can contain at most one such mapping.)
	*
	* <p>Returns the value to which this map previously associated the key,
	* or {@code null} if the map contained no mapping for the key.
	*
	* <p>If this map permits null values, then a return value of
	* {@code null} does not <i>necessarily</i> indicate that the map
	* contained no mapping for the key; it's also possible that the map
	* explicitly mapped the key to {@code null}.
	*
	* <p>The map will not contain a mapping for the specified key once the
	* call returns.
	*
	* @param key key whose mapping is to be removed from the map
	* @return the previous value associated with {@code key}, or
	*         {@code null} if there was no mapping for {@code key}.
	* @throws UnsupportedOperationException if the {@code remove} operation
	*         is not supported by this map
	* @throws ClassCastException if the key is of an inappropriate type for
	*         this map
	* (<a href="{@docRoot}/java.base/java/util/Collection.html#optional-restrictions">optional</a>)
	* @throws NullPointerException if the specified key is null and this
	*         map does not permit null keys
	* (<a href="{@docRoot}/java.base/java/util/Collection.html#optional-restrictions">optional</a>)
	 */
	Remove(key basic.Object) V

	// Bulk Operations

	/**
	* Copies all of the mappings from the specified map to this map
	* (optional operation).  The effect of this call is equivalent to that
	* of calling {@link #put(Object,Object) put(k, v)} on this map once
	* for each mapping from key {@code k} to value {@code v} in the
	* specified map.  The behavior of this operation is undefined if the
	* specified map is modified while the operation is in progress.
	*
	* @param m mappings to be stored in this map
	* @throws UnsupportedOperationException if the {@code putAll} operation
	*         is not supported by this map
	* @throws ClassCastException if the class of a key or value in the
	*         specified map prevents it from being stored in this map
	* @throws NullPointerException if the specified map is null, or if
	*         this map does not permit null keys or values, and the
	*         specified map contains null keys or values
	* @throws IllegalArgumentException if some property of a key or value in
	*         the specified map prevents it from being stored in this map
	 */
	// void putAll(HashMap<? extends K, ? extends V> m);
	PutAll(m HashMap[K, V])

	/**
	* Removes all of the mappings from this map (optional operation).
	* The map will be empty after this call returns.
	*
	* @throws UnsupportedOperationException if the {@code clear} operation
	*         is not supported by this map
	 */
	Clear()

	/**
	* Returns a {@link Set} view of the keys contained in this map.
	* The set is backed by the map, so changes to the map are
	* reflected in the set, and vice-versa.  If the map is modified
	* while an iteration over the set is in progress (except through
	* the iterator's own {@code remove} operation), the results of
	* the iteration are undefined.  The set supports element removal,
	* which removes the corresponding mapping from the map, via the
	* {@code Iterator.remove}, {@code Set.remove},
	* {@code removeAll}, {@code retainAll}, and {@code clear}
	* operations.  It does not support the {@code add} or {@code addAll}
	* operations.
	*
	* @return a set view of the keys contained in this map
	 */
	// KeySet() *util.Set[K]

	// Comparison and hashing

	/**
	* Compares the specified object with this map for equality.  Returns
	* {@code true} if the given object is also a map and the two maps
	* represent the same mappings.  More formally, two maps {@code m1} and
	* {@code m2} represent the same mappings if
	* {@code m1.entrySet().equals(m2.entrySet())}.  This ensures that the
	* {@code equals} method works properly across different implementations
	* of the {@code HashMap} interface.
	*
	* @param o object to be compared for equality with this map
	* @return {@code true} if the specified object is equal to this map
	 */
	Equals(o basic.Object) bool

	/**
	* Returns the hash code value for this map.  The hash code of a map is
	* defined to be the sum of the hash codes of each entry in the map's
	* {@code entrySet()} view.  This ensures that {@code m1.equals(m2)}
	* implies that {@code m1.hashCode()==m2.hashCode()} for any two maps
	* {@code m1} and {@code m2}, as required by the general contract of
	* {@link Object#hashCode}.
	*
	* @return the hash code value for this map
	* @see HashMap.Entry#hashCode()
	* @see Object#equals(Object)
	* @see #equals(Object)
	 */
	HashCode() int32
}
