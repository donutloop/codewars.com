def convert_hash_to_array(hash)
 res = [] of Array(String|Int32)
 hash.each do |key, value|

      element = [] of String|Int32
      element << key
      element << value
      res << element
 end

 res = res.sort { |x, y| x[0].to_s <=> y[0].to_s }
 return res
end