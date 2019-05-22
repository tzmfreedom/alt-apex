#!/usr/bin/env ruby

node_types = %w[
  File
  Header
  Class
  Property
  Constructor
  Parameter
  Method
  Modifier
  If
  For
  Switch
  When
  VariableDeclaration
  Block
  TypeRef
  Boolean
  Integer
  String
  Identifier
  MemberAccess
  BinaryOperator
  Name
]

=begin
puts 'type Visitor interface {'
node_types.each do |node_type|
  puts "\tVisit#{node_type}(*#{node_type}) (interface{}, error)"
end
puts '}'

node_types.each do |node_type|
  puts <<EOS
func (t *#{node_type}) Accept(v Visitor) (interface{}, error) {
\treturn v.Visit#{node_type}(t)
}

EOS
end
=end

name = 'StringVisitor'
puts "type #{name} struct {}"
node_types.each do |node_type|
  puts <<EOS
func (v *#{name}) Visit#{node_type}(n *#{node_type}) (interface{}, error) {
}

EOS
end
